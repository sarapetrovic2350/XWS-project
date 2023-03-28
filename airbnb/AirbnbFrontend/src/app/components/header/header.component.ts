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
  isAdministrator: boolean = false;
  isRegisteredUser: boolean = false;

  constructor(private router: Router, private userService : UserService) { }

  ngOnInit(): void {
    this.refreshUser()
  }

  refreshUser(): void {
    let user = this.userService.getCurrentUser()
    console.log(user)
    if(user === null) {
      this.isLoggedIn = false;
    } else {
      this.isLoggedIn = true;
      if(user.role == "ADMINISTRATOR") {
        this.isAdministrator = true;
      }
      if(user.role == "REGISTERED_USER") {
        this.isRegisteredUser = true;
      }
    }

  }

  logout(){
    this.userService.logout();
  }
}
