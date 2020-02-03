import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgZorroAntdModule, NZ_I18N, en_US } from 'ng-zorro-antd';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { registerLocaleData } from '@angular/common';
import en from '@angular/common/locales/en';
import { LinkModule } from './link/link.module';
import { JwtInterceptor } from './auth/jwt.interceptor';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { AuthModule } from './auth/auth.module';
import { ErrorInterceptor } from './auth/error.interceptor';
import { EnvInterceptor } from './auth/env.interceptor';

registerLocaleData(en);

@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgZorroAntdModule,
    FormsModule,
    BrowserAnimationsModule,
    LinkModule,
    AuthModule
  ],
  providers: [
    { provide: NZ_I18N, useValue: en_US },
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: ErrorInterceptor, multi: true },
    { provide: HTTP_INTERCEPTORS, useClass: EnvInterceptor, multi: true },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
