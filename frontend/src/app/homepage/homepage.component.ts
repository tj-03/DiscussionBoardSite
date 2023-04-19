import { HttpClient, HttpClientModule, HttpHeaders } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgModel } from '@angular/forms';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css'],
  providers: [HttpClient],
})

export class HomepageComponent implements OnInit {
  posts: any;
  constructor(private httpClient: HttpClient) {
    this.posts = [];
  }

  
  ngOnInit() {
    let url = 'http://localhost:8080/api';
    let header = new HttpHeaders();
    this.httpClient.get('http://localhost:8080/api/posts', {
      headers: header
    }).subscribe((posts : Object) => {
      this.posts = posts;
      for(let i = 0; i < this.posts.length; i++){
        this.httpClient.get('http://localhost:8080/api/user/' + this.posts[i].author_id, {
          headers: header
        })
        .subscribe(
          {
            next: (user: any) => {
              this.posts[i].author = user;
            },
            error: (err: any) => {
              console.error(err);
              this.posts[i].author = {name: 'unknown user'};
            }
          }
        );

      }
    });
  }


}
