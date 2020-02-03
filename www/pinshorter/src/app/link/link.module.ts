import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateLinkComponent } from './create-link/create-link.component';
import { NgZorroAntdModule } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { ListLinkComponent } from './list-link/list-link.component';
import { RouterModule } from '@angular/router';



@NgModule({
  declarations: [CreateLinkComponent, ListLinkComponent],
  imports: [
    CommonModule,
    NgZorroAntdModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    RouterModule,
  ],
  exports: [
    CreateLinkComponent,
    ListLinkComponent
  ]
})
export class LinkModule { }
