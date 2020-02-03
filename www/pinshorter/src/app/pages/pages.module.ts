import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LinksComponent } from './links/links.component';
import { LinkModule } from '../link/link.module';



@NgModule({
  declarations: [LinksComponent],
  imports: [
    CommonModule,
    LinkModule
  ],
  exports: [
    LinksComponent
  ]
})
export class PagesModule { }
