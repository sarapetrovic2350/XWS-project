import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {LoginRequest, User} from "../model/user.model";

@Injectable({
  providedIn: 'root'
})
export class UserService {
  apiHost: string = 'http://localhost:8080/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json'});
  private user = new User();

  constructor(private http: HttpClient) { }

  registerUser(user: User) {
    return this.http.post<User>(this.apiHost + 'users/', user);
  }

  login(request: LoginRequest){
    return this.http.post<any>(this.apiHost + 'users/login', request);
  }

  setLoggedUser(data: any) {
    this.user = data;
    console.log(this.user)
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
