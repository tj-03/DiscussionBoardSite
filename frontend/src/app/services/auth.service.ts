import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {GoogleAuthProvider, signInWithPopup} from "@angular/fire/auth";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import { Router } from '@angular/router';


interface UserData {
  userId: string | null;
  userName: string | null;
  userToken: string | null;
}

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  

  userData: UserData = {
    userId: null,
    userName: null,
    userToken: null
  }
  constructor(public authService: AngularFireAuth, private router: Router) {
    let emptyUserData = `
      {
        "userId": null,
        "userName": null,
        "userToken": null
      }
    `
    this.userData = JSON.parse(localStorage.getItem('userData') || emptyUserData);
    this.authService.authState.subscribe(user => {
      user?.getIdToken().then((token) => {
        if (token) {
          console.log(token)
          this.userData = {
            userId: user?.uid,
            userName: user?.email,
            userToken: token
          }
          localStorage.setItem('userData', JSON.stringify(this.userData.userToken));
        } else {
          localStorage.setItem('userData', '');
        }
      });
   })
  }

  login(){
    return this.authService
      .signInWithPopup(new GoogleAuthProvider())
      .then((userCredential) => {   
        this.userData.userId = userCredential.user?.uid!;
        this.userData.userName = userCredential.user?.email!;
        return userCredential.user?.getIdToken();
      })
      .then((token) => {
        if (token) {
          this.userData.userToken = token;
          localStorage.setItem('userToken', JSON.stringify(this.userData));
        } else {
          localStorage.setItem('userToken', '');
        }
      })
      .catch((error) => {
        window.alert(error);
      });
  }

  logout(){
    return this.authService.signOut().then(() => {
      localStorage.removeItem('userData');
      this.userData = {
        userId: null,
        userName: null,
        userToken: null
      }
    });
  }
}
