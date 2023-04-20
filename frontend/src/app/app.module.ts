import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatButtonModule} from '@angular/material/button';
import { AboutteamComponent } from './aboutteam/aboutteam.component';
import { CreatepostComponent } from './createpost/createpost.component';
import {MatInputModule} from '@angular/material/input';
import { LoginComponent } from './login/login.component';
import { DialogContentExampleDialog, HomepageComponent } from './homepage/homepage.component';
import { StepperComponent } from './stepper/stepper.component';
import {MatCardModule} from '@angular/material/card';
import {MatChipsModule} from '@angular/material/chips';
import {MatDividerModule} from '@angular/material/divider';
import {MatBadgeModule} from '@angular/material/badge';
import {MatIconModule} from '@angular/material/icon';
import { SignupComponent } from './signup/signup.component';
import { HttpClientModule } from '@angular/common/http';
import {AngularFireModule} from '@angular/fire/compat';
import { AngularFireAuthModule } from '@angular/fire/compat/auth';
import { FormsModule } from '@angular/forms';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';

const firebaseConfig = {
  apiKey: "AIzaSyAHzbLjsiQswNjy45AAqsG0MHJU-iibdl4",
  authDomain: "groupproject-333e4.firebaseapp.com",
  projectId: "groupproject-333e4",
  storageBucket: "groupproject-333e4.appspot.com",
  messagingSenderId: "345612458789",
  appId: "1:345612458789:web:40d1daf0c95280400dd8cc",
  measurementId: "G-PKJVPSEBT6"
};
@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    AboutteamComponent,
    CreatepostComponent,
    LoginComponent,
    HomepageComponent,
    StepperComponent,
    SignupComponent,
    DialogContentExampleDialog
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MatToolbarModule,
    MatButtonModule,
    BrowserAnimationsModule,
    MatInputModule,
    MatChipsModule,
    MatCardModule,
    MatDividerModule,
    MatBadgeModule,
    MatIconModule,
    MatDialogModule,
    HttpClientModule,
    AngularFireModule.initializeApp(firebaseConfig),
    AngularFireAuthModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
