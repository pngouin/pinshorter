import { Component, OnInit } from '@angular/core';
import { LinkService } from '../link.service';
import { Link } from '../link.model';

@Component({
  selector: 'app-list-link',
  templateUrl: './list-link.component.html',
  styleUrls: ['./list-link.component.scss']
})
export class ListLinkComponent implements OnInit {

  constructor(private linkService: LinkService) { }

  ngOnInit() {
    this.linkService.list().subscribe(links => {
      this.listOfData = links;
      this.displayedData = [...this.listOfData]
    })
  }

  listOfData: Link[] = [];
  displayedData: Link[] = [];

  delete(link: Link): void {
    this.linkService.delete(link).subscribe(
      () => {
        const index = this.listOfData.indexOf(link);
        if (index > -1) {
          this.listOfData.splice(index, 1);
          this.displayedData = [...this.listOfData]
        }
      }
    );
  }
}
