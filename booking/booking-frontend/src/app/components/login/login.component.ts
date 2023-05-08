import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoginRequest } from 'src/app/model/user.model';
import { UserService } from 'src/app/service/user.service';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private router: Router,
    private userService: UserService) {

   }
  title = 'Login to Airbnb';
  message: string= "";
  request = new LoginRequest();
  submitted = false;

  ngOnInit(): void {
  }

  onSubmit() {
    if (this.request.email == '' || this.request.password == '') {
      this.message = 'Email or password missing.';
    } else {
      this.submitted = true;
      this.userService.login(this.request).subscribe(
        {
          next: (res) => {
            console.log(res)
            this.successfulLogin(res);
            Swal.fire({
              icon: 'success',
              title: 'Success!',
              text: 'Sucessfully logged in!',
            })
            window.location.href = '/';

          },
          error: (e) => {
            this.submitted = false;
            console.log(e);
            Swal.fire({
              icon: 'error',
              title: 'Oops...',
              text: 'Invalid email or password.',
            })

          }
        });

    }
  }
  successfulLogin(token: string) {
    this.userService.setTokenForLoggedInUser(token);
  }
}
