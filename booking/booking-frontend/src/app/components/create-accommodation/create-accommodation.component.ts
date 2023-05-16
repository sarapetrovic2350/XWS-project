import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";
import Swal from 'sweetalert2';
import { Accommodation } from 'src/app/model/accommodation.model';
import { AccommodationService } from 'src/app/service/accommodation.service';
import {UserService} from "../../service/user.service";
import {User} from "../../model/user.model";
import {FormControl, Validators, FormGroup, FormBuilder} from '@angular/forms';

@Component({
  selector: 'app-create-accommodation',
  templateUrl: './create-accommodation.component.html',
  styleUrls: ['./create-accommodation.component.css']
})
export class CreateAccommodationComponent implements OnInit {

  constructor(
    private router: Router,
    private accommodationService: AccommodationService,
    private userService : UserService
  ) { }

  title = 'Create an Accommodation';
  accommodation = new Accommodation();
  submitted = false;
  street: string = "";
  streetNumber: string = "";
  city: string= "";
  country: string="";
  benefits: string[] = [];
  auxT: any[] = [];

  BenefitList: string[] = ['Wifi', 'Free Parking', 'Private Bathroom', 'Shared Bathroom', 'Kitchen',
  'Air Conditioner', 'Kitchen'];


  isLoggedIn: boolean = false;
  isHost: boolean = false;
  isGuest: boolean = false;
  user: User = new User();
  ngOnInit(): void {
    let userRole = this.userService.getLoggedInUserRole()
    // let userEmail = this.userService.getLoggedInUserEmail()
    // this.userService.getUserByEmail(userEmail).subscribe(res => {
    //   this.user = res; 
    //   console.log(this.user)
    // })

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

  onSubmit(){
    
    let userId = this.userService.getLoggedInUserId()

    console.log(userId)
    this.accommodation.address.street = this.street;
    this.accommodation.address.number = this.streetNumber;
    this.accommodation.address.city = this.city;
    this.accommodation.address.country = this.country;
    if (this.benefits != null) {
      for (let t of this.benefits) {
        this.auxT.push(t);
      }
    }

    
    this.accommodation.benefits = this.auxT;
    //this.accommodation.hostId = userid;
    console.log(this.accommodation.hostId);
    var NewAccommodation = {
      name: this.accommodation.name,
      minNumberOfGuests: this.accommodation.minNumberOfGuests,
      maxNumberOfGuests: this.accommodation.maxNumberOfGuests,
      address: this.accommodation.address,
      hostID: userId, 
      benefits: this.accommodation.benefits
    }
    console.log(NewAccommodation);

    this.accommodationService.createAccommodation(NewAccommodation).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/show-host-accommodations']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully created new accommodation!',
          })

        },
        error: (e) => {
          this.submitted = false;
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong.',
          })

        }

      });

  }

}
