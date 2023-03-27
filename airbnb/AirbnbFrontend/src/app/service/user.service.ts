import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {LoginRequest, User} from "../model/user.model";

@Injectable({
  providedIn: 'root'
})
export class UserService {
  apiHost: string = 'http://localhost:8080/';
  headers: HttpHeaders = new HttpHeaders({ 'Content-Type': 'application/json'});
  constructor(private http: HttpClient) { }
  registerUser(user: User) {
    return this.http.post<User>(this.apiHost + 'users/', user);
  }

  login(request: LoginRequest){
    return this.http.post<any>(this.apiHost + 'users/login', request);
  }
}