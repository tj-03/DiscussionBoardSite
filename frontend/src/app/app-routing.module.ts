import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutteamComponent } from './aboutteam/aboutteam.component';
import { CreatepostComponent } from './createpost/createpost.component';
import { HomepageComponent } from './homepage/homepage.component';
import { LoginComponent } from './login/login.component';

const routes: Routes = [
  { path: 'about', component: AboutteamComponent },
  { path: 'createpost', component: CreatepostComponent },
  { path: 'login', component: LoginComponent },
  { path: 'homepage', component: HomepageComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
