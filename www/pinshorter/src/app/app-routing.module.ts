import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateLinkComponent } from './link/create-link/create-link.component';
import { AuthGuard } from './auth/auth.guard';
import { LoginComponent } from './auth/login/login.component';


const routes: Routes = [
  { path: '', component: CreateLinkComponent, canActivate: [AuthGuard] },
  { path: 'login', component: LoginComponent },
  { path: '*', redirectTo: '' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
