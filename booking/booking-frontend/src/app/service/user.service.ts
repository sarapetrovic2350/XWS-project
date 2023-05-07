import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {LoginRequest, User} from "../model/user.model";

@Injectable({
  providedIn: 'root'
})
export class UserService {
  apiHost: string = 'http://localhost:8080/api/user/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json', 'Access-Control-Allow-Origin' : '*',
    'Access-Control-Allow-Methods' : 'GET,HEAD,OPTIONS,POST,PUT', 'Access-Control-Allow-Headers' : 'Origin, X-Requested-With, Content-Type, Accept, x-client-key, x-client-token, x-client-secret, Authorization' }  );
  private user = new User();

  constructor(private http: HttpClient) { }

  registerUser(user: User) {
    return this.http.post<User>(this.apiHost, user);
  }

  login(request: LoginRequest){
    return this.http.post<any>(this.apiHost + 'login', request);
  }

  setLoggedUser(data: any) {
    this.user = data;
    console.log(this.user)
    console.log(this.user.role)
    localStorage.setItem('currentUser', JSON.stringify(this.user));
    console.log(localStorage.getItem('currentUser'));
  }

  getCurrentUser(): User {
    return JSON.parse(localStorage.getItem('currentUser')!);
  }

  logout(){
    localStorage.clear();
    window.location.href = 'login';
  }
}
