import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LinksComponent } from './links/links.component';
import { LinkModule } from '../link/link.module';
import { NgZorroAntdModule } from 'ng-zorro-antd';
import { LoginComponent } from './login/login.component';
import { AuthModule } from '../auth/auth.module';



@NgModule({
  declarations: [LinksComponent, LoginComponent],
  imports: [
    CommonModule,
    LinkModule,
    NgZorroAntdModule,
    AuthModule
  ],
  exports: [
    LinksComponent
  ]
})
export class PagesModule { }
