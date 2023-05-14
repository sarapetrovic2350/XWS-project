import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import {UserService} from "../../service/user.service";
import {User} from "../../model/user.model";

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  isLoggedIn: boolean = false;
  isHost: boolean = false;
  isGuest: boolean = false;
  user: User = new User();

  constructor(private router: Router, private userService : UserService) { }

  ngOnInit(): void {
    this.refreshUser()
  }

  refreshUser(): void {
    let userRole = this.userService.getLoggedInUserRole()
    let userEmail = this.userService.getLoggedInUserEmail()
    console.log(userEmail)
    this.userService.getUserByEmail(userEmail).subscribe(res => {
      this.user = res;
      console.log(this.user)
    })

    if(userRole === "") {
      this.isLoggedIn = false;
    } else {
      this.isLoggedIn = true;
      if(userRole == "HOST") {
        this.isHost = true;
      }
      if(userRole == "GUEST") {
        this.isGuest = true;
      }
    }

  }

  logout(){
    this.userService.logout();
  }
}
