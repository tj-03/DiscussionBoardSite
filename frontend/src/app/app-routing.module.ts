import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutteamComponent } from './aboutteam/aboutteam.component';
import { CreatepostComponent } from './createpost/createpost.component';

const routes: Routes = [
  { path: 'about', component: AboutteamComponent },
  { path: 'createpost', component: CreatepostComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
