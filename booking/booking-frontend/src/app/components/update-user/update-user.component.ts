import { Component, OnInit } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { User } from 'src/app/model/user.model';
import { UserService } from 'src/app/service/user.service';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-update-user',
  templateUrl: './update-user.component.html',
  styleUrls: ['./update-user.component.css']
})
export class UpdateUserComponent implements OnInit {

  constructor(private router: Router, private userService : UserService) { }

  title = 'Update informations';
  user = new User()
  passwordRepeated: string= "";
  submitted = false;

  ngOnInit(): void {
    let userRole = this.userService.getLoggedInUserRole()
    let userEmail = this.userService.getLoggedInUserEmail()
    let userId = this.userService.getLoggedInUserId()
    this.userService.getUserByEmail(userEmail).subscribe(res => {
      this.user = res.user;
      console.log(res)
    })
  }

  email = new FormControl('', [Validators.required, Validators.email]);
  name = new FormControl('', [Validators.required, Validators.minLength(2), Validators.maxLength(30)])
  lastName = new FormControl('', [Validators.required, Validators.minLength(2), Validators.maxLength(30)])
  password = new FormControl('', [Validators.required, Validators.minLength(6), Validators.maxLength(30)])
  passwordConfirm = new FormControl('', [Validators.required, Validators.minLength(6), Validators.maxLength(30)])
  streetFormControl = new FormControl('', [Validators.required, Validators.minLength(2), Validators.maxLength(30)])
  streetNumberFormControl = new FormControl('', [Validators.required])
  cityFormControl = new FormControl('', [Validators.required, Validators.minLength(2), Validators.maxLength(30)])
  countryFormControl = new FormControl('', [Validators.required, Validators.minLength(2), Validators.maxLength(30)])
  roleFormControl = new FormControl( '', [Validators.required])
  getEmailErrorMessage() {
    return this.email.hasError('required') ? 'You must enter email' :
      this.email.hasError('email') ? 'Not a valid email' :
        '';
  }
  passwordMatchValidator() {
    return this.user.password === this.passwordRepeated
  }
  onSubmit(){
    this.userService.updateUser(this.user).subscribe( 
      {next: (res) => {
        this.router.navigate(['/update-user']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Sucessfully updated!',
          })  
      },
      error: (e) => {
        console.log(e);
        this.submitted = false;
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Check the fields again!',
          })  
      }
      });

  }

  deleteAccount(){
    let userId = this.userService.getLoggedInUserId()
      this.userService.deleteAccount(userId).subscribe(
        {
          next: (res) => {
            localStorage.clear()
            //this.router.navigate(['/register-user']);
            window.location.href = 'register-user';
            Swal.fire({
              icon: 'success',
              title: 'Success!',
              text: 'Successfully deleted account!',
            })
  
          },
          error: (e) => {
            this.submitted = false;
            console.log(e);
            Swal.fire({
              icon: 'error',
              title: 'Oops...',
              text: 'You have active reservations.',
            })
  
          }
  
        });
  }

}
