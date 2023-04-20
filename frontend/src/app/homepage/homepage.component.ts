import { HttpClient, HttpClientModule, HttpHeaders } from '@angular/common/http';
import { Component, EventEmitter, Inject, Input, OnInit, Output } from '@angular/core';
import { NgModel } from '@angular/forms';
import { AuthService } from '../services/auth.service';
import {MatDialog, MatDialogRef, MAT_DIALOG_DATA} from '@angular/material/dialog';



@Component({
  selector: 'dialog-content-example-dialog',
  templateUrl: 'dialog.component.html',
  styleUrls: ['./homepage.component.css']
})
export class DialogContentExampleDialog {
  

  comments: any = [];
  commentContent: string = '';
  constructor(
    public dialogRef: MatDialogRef<DialogContentExampleDialog>,
    @Inject(MAT_DIALOG_DATA) public post: any,
    private httpClient: HttpClient,
    private auth: AuthService) { 
      console.log(post)
    }
  
  onCancel(): void {
    this.dialogRef.close();
  }


  ngOnInit() {
    this.httpClient.get('http://localhost:8080/api/post/comments/' + this.post.id, {
    }).subscribe((comments: any) => {
      this.comments = comments || [];
      console.log(this.comments);
      for(let comment of comments){
        this.httpClient.get('http://localhost:8080/api/user/' + comment.author_id, {
        }).subscribe((user: any) => {
          if(user)
          comment.author = user;
        });
      }
    });
  }

  
  submitComment() {
    let headers = new HttpHeaders().set('Authorization', this.auth.userData.userToken!);
    this.httpClient.post('http://localhost:8080/api/post/comment', {
      post_id: this.post.id,
      content: this.commentContent,
      author_id: this.auth.userData.userId!
    },{
      headers: headers
    }).subscribe((res: any) => {
      console.log(res);
      let newComments = [...this.comments];
      newComments.push({
        content: this.commentContent,
        author:{
          name: this.auth.userData.userName
        }
      });
      this.comments = newComments;
    });

  }
  
}

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css'],
  providers: [HttpClient],
})
export class HomepageComponent implements OnInit {
  posts: any;
  constructor(private httpClient: HttpClient, public dialog: MatDialog, private auth:AuthService) {
    this.posts = [];
  }

  openDialog(index: any){
    let dialogRef = this.dialog.open(DialogContentExampleDialog, {
      width: '90vw',
      data:this.posts[index],
      backdropClass: 'backdropClass'
    });
  }

  ngOnInit() {
    let url = 'http://localhost:8080/api';
    let header = new HttpHeaders();
    this.httpClient.get('http://localhost:8080/api/posts', {
      headers: header
    }).subscribe((posts : Object) => {
      this.posts = posts;
      for(let i = 0; i < this.posts.length; i++){
        if (!this.posts[i].title){
          this.posts[i].title = 'Untitled';
        }
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
