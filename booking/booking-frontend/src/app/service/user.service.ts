import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {LoginRequest, User} from "../model/user.model";
import jwt_decode from "jwt-decode";
import {Observable} from "rxjs";
import {Flight} from "../model/flight.model";

@Injectable({
  providedIn: 'root'
})
export class UserService {
  apiHost: string = 'http://localhost:8000/user';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Access-Control-Allow-Origin' : '*',
    'Access-Control-Allow-Methods' : 'GET,HEAD,OPTIONS,POST,PUT', 'Access-Control-Allow-Headers' : 'Origin, X-Requested-With, Content-Type, Accept, x-client-key, x-client-token, x-client-secret, Authorization' }  );
  headers2: HttpHeaders = new HttpHeaders({
    'Content-Type': 'application/json',
  });
  private user = new User();
  token: string = ""

  constructor(private http: HttpClient) { }

  registerUser(user: User): Observable<any> {
    return this.http.post<User>(this.apiHost, user, {headers: this.headers2});
  }

  login(request: LoginRequest){
    return this.http.post<any>(this.apiHost + '/login', request);
  }
  getDecodedAccessToken(token: string): any {
    try {
      return jwt_decode(token);
    } catch(Error) {
      return null;
    }
  }
  setTokenForLoggedInUser(token: string) {
    localStorage.setItem('token', token);
    console.log(localStorage.getItem('token'));
  }

  getLoggedInUserRole(): string {
    const decodedToken = this.getDecodedAccessToken(localStorage.getItem("token")!);
    if (decodedToken == null) {
      return "";
    } else {
      return decodedToken.role
    }
  }

  getLoggedInUserEmail(): string {
    const decodedToken = this.getDecodedAccessToken(localStorage.getItem("token")!);
    if (decodedToken == null) {
      return "";
    } else {
      return decodedToken.email
    }
  }

  getLoggedInUserId(): string {
    const decodedToken = this.getDecodedAccessToken(localStorage.getItem("token")!);
    if (decodedToken == null) {
      return "";
    } else {
      return decodedToken.id
    }
  }

  getUserByEmail(email: string): Observable<any> {
    return this.http.get<any>(this.apiHost + '/' + email);
  }

  updateUser(user: User): Observable<any> {
    console.log(user);
    return this.http.post<User>(this.apiHost + '/update', user, {headers: this.headers2});
  }


  logout(){
    localStorage.clear();
    window.location.href = 'login';
  }
}
