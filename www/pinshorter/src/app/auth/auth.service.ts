import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Connection, Token, User } from './auth.model';
import { Observable, BehaviorSubject } from 'rxjs';
import { map } from 'rxjs/operators'
import { JsonPipe } from '@angular/common';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient) {
    this.currentUserSubject = new BehaviorSubject<User>(JSON.parse(localStorage.getItem('user')));
    this.currentUser = this.currentUserSubject.asObservable();
   }

  private currentUserSubject: BehaviorSubject<User>;
  public currentUser: Observable<User>;

  auth(conn: Connection): Observable<Token> {
    return this.http.post<Token>(`/auth`, conn).pipe(map(
        token => {
          let user = this.jwtInformation(token.token)
          localStorage.setItem('user', JSON.stringify(user));
          this.currentUserSubject.next(user);
          return token;
        }
      ))
  }

  public get currentUserValue(): User {
    return this.currentUserSubject.value;
  }

  logout() {
    localStorage.removeItem('user');
    this.currentUserSubject.next(null);
  }

  jwtInformation(str: string): User {
    let user = JSON.parse(atob(str.split('.', -1)[1])) as User;
    user.token = str;
    return user;
  }
  
}
