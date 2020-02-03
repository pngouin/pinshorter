import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthGuard } from './auth/auth.guard';
import { LoginComponent } from './auth/login/login.component';
import { LinksComponent } from './pages/links/links.component';


const routes: Routes = [
  { path: '', component: LinksComponent, canActivate: [AuthGuard] },
  { path: 'login', component: LoginComponent },
  { path: '*', redirectTo: '' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
