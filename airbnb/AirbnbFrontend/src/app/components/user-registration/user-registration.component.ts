import { Component, OnInit } from '@angular/core';
import {UserService} from "../../service/user.service";
import {Router} from "@angular/router";
import {User} from "../../model/user.model";
import Swal from 'sweetalert2';

@Component({
  selector: 'app-user-registration',
  templateUrl: './user-registration.component.html',
  styleUrls: ['./user-registration.component.css']
})
export class UserRegistrationComponent implements OnInit {

  constructor(
    private router: Router,
    private userService: UserService
  ) { }
  title = 'Register to Airbnb';
  user = new User();
  submitted = false;
  passwordRepeated: string= "";
  street: string = "";
  streetNumber: string = "";
  city: string= "";
  country: string="";

  ngOnInit(): void {
  }

  onSubmit() {
    this.user.address.street = this.street;
    this.user.address.streetNumber = this.streetNumber;
    this.user.address.city = this.city;
    this.user.address.country = this.country;

    this.userService.registerUser(this.user).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/login']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully registered to Airbnb!',
          })

        },
        error: (e) => {
          this.submitted = false;
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Email already exists.',
          })

        }

      });
  }
}
