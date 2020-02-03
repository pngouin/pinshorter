import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Link } from './link.model';
import { Observable } from 'rxjs';



@Injectable({
  providedIn: 'root'
})
export class LinkService {

  constructor(private httpClient: HttpClient) { }

  add(link: Link): Observable<Link> {
    return this.httpClient.put<Link>(`/link`, link);
  }

  delete(link: Link): Observable<void> {
    return this.httpClient.delete<void>(`/link/${link.apiPoint}`);
  }

  list(): Observable<Link[]> {
    return this.httpClient.get<Link[]>(`/link`);
  }

}
