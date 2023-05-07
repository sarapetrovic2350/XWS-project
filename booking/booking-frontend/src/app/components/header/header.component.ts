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
  isHost: boolean = false;
  isGuest: boolean = false;

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
      if(user.role == "HOST") {
        this.isHost = true;
      }
      if(user.role == "GUEST") {
        this.isGuest = true;
      }
    }

  }

  logout(){
    this.userService.logout();
  }
}
