import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LinkService } from '../link.service';
import { Link } from '../link.model';

@Component({
  selector: 'app-create-link',
  templateUrl: './create-link.component.html',
  styleUrls: ['./create-link.component.scss']
})
export class CreateLinkComponent implements OnInit {

  validateForm: FormGroup;
  @Output() created = new EventEmitter<Link>();

  constructor(private fb: FormBuilder, private linkService: LinkService) { }

  submitForm(): void {
    for (const i in this.validateForm.controls) {
      this.validateForm.controls[i].markAsDirty();
      this.validateForm.controls[i].updateValueAndValidity();
    }
    this.linkService.add(this.validateForm.value as Link).subscribe(link =>  {
      this.created.emit(link)
      this.copyToClipboard(link.api_point)
    });
    this.validateForm.reset();
  }

  copyToClipboard(str: string) {
    document.addEventListener('copy', (e: ClipboardEvent) => {
      e.clipboardData.setData('text/plain', (`${window.location.origin}/${str}`));
      e.preventDefault();
      document.removeEventListener('copy', null);
    });
    document.execCommand('copy');
  }

  ngOnInit() {
    this.validateForm = this.fb.group({
      url: [null, [Validators.required]],
    });
  }

}
