import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateLinkComponent } from './create-link/create-link.component';
import { NgZorroAntdModule } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';



@NgModule({
  declarations: [CreateLinkComponent],
  imports: [
    CommonModule,
    NgZorroAntdModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
  ],
  exports: [
    CreateLinkComponent
  ]
})
export class LinkModule { }
