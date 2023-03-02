import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { NgModel } from '@angular/forms';


@Component({
  selector: 'homepage',
  templateUrl: './homepage.component.html',
  providers: [HttpClient],
})

export class HomepageComponent implements OnInit {
  posts: any;
  constructor(private httpClient: HttpClient) {
    this.posts = [];
  }

  ngOnInit() {
    this.httpClient.get('http://localhost:8080/api/posts').subscribe((posts : Object) => {
      this.posts = posts;
      console.log(this.posts);
    });
  }
}