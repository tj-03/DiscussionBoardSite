import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, EventEmitter, Output } from '@angular/core';
import { AuthService } from '../services/auth.service';

@Component({
  selector: 'app-createpost',
  templateUrl: './createpost.component.html',
  styleUrls: ['./createpost.component.css']
})
export class CreatepostComponent {
  postContent: string = '';
  constructor(private httpClient: HttpClient, private auth: AuthService) { 
    //console.log(this.auth.userData);
  }

  submitPost() {
    let header = new HttpHeaders();
    let token = this.auth.userData.userToken!;
    header = header.set('Authorization', token);

    let post = {
      author_id: this.auth.userData.userId!,
      content: this.postContent
    }

    this.httpClient.post('http://localhost:8080/api/post', post, {
      headers: header
    }).subscribe((res : Object) => {
      console.log(res);
    });
  }
}
