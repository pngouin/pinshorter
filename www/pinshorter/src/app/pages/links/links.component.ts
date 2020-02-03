import { Component, OnInit } from '@angular/core';
import { Link } from 'src/app/link/link.model';
import { Subject } from 'rxjs';

@Component({
  selector: 'app-links',
  templateUrl: './links.component.html',
  styleUrls: ['./links.component.scss']
})
export class LinksComponent implements OnInit {

  constructor() { }

  parentSubject:Subject<Link> = new Subject();

  ngOnInit() {
  }

  linkCreated(link: Link){
    this.parentSubject.next(link);
  }

}
