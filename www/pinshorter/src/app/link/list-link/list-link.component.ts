import { Component, OnInit, Input } from '@angular/core';
import { LinkService } from '../link.service';
import { Link } from '../link.model';
import { Subject } from 'rxjs';

@Component({
  selector: 'app-list-link',
  templateUrl: './list-link.component.html',
  styleUrls: ['./list-link.component.scss']
})
export class ListLinkComponent implements OnInit {

  constructor(private linkService: LinkService) { }
  
  @Input() listener:Subject<Link>;

  ngOnInit() {
    this.listener.subscribe( link => {
      this.listOfData.push(link)
      this.displayedData = [...this.listOfData]
    })
    this.linkService.list().subscribe(links => {
      this.listOfData = links;
      this.displayedData = [...this.listOfData]
    })
  }

  listOfData: Link[] = [];
  displayedData: Link[] = [];

  getApiURL(link: Link): string {
    return `${window.location.origin}/${link.api_point}`
  }

  delete(link: Link): void {
    this.linkService.delete(link).subscribe(
      () => {
        const index = this.listOfData.indexOf(link);
        if (index > -1) {
          console.log(link)
          this.listOfData.splice(index, 1);
          this.displayedData = [...this.listOfData]
        }
      }
    );
  }
}
