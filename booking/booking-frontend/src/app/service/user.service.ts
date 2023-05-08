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
  apiHost: string = 'http://localhost:8080/api/user/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Access-Control-Allow-Origin' : '*',
    'Access-Control-Allow-Methods' : 'GET,HEAD,OPTIONS,POST,PUT', 'Access-Control-Allow-Headers' : 'Origin, X-Requested-With, Content-Type, Accept, x-client-key, x-client-token, x-client-secret, Authorization' }  );
  private user = new User();
  token: string = ""

  constructor(private http: HttpClient) { }

  registerUser(user: User) {
    return this.http.post<User>(this.apiHost, user);
  }

  login(request: LoginRequest){
    return this.http.post<any>(this.apiHost + 'login', request);
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
  getUserByEmail(email: string): Observable<User> {
    return this.http.get<User>(this.apiHost + 'userByEmail/' + email);
  }


  logout(){
    localStorage.clear();
    window.location.href = 'login';
  }
}