import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {UserService} from "../../service/user.service";

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  isLoggedIn: boolean = false;

  constructor(private router: Router, private userService : UserService) { }

  ngOnInit(): void {
    this.refreshUser()
  }

  signUp() {
    this.router.navigate(['register-user']);
  }
  home() {
    this.router.navigate(['/']);
  }
  login() {
    this.router.navigate(['login']);
  }

  refreshUser(): void {
    let user = localStorage.getItem('currentUser');
    this.userService.getCurrentUser()
    console.log(user)
    if(user === null ){
      this.isLoggedIn = false;
    } else{
      this.isLoggedIn = true;
    }
  }

  logout(){
    this.userService.logout();
  }
}
