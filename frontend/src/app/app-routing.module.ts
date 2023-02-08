import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutteamComponent } from './aboutteam/aboutteam.component';

const routes: Routes = [
  { path: 'about', component: AboutteamComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
