import { Component, OnInit } from '@angular/core';
import {Router, ActivatedRoute, Params} from "@angular/router";
import Swal from 'sweetalert2';
import { AccommodationService } from 'src/app/service/accommodation.service';
import { User } from 'src/app/model/user.model';
import { UserService } from 'src/app/service/user.service';
import { Availability } from 'src/app/model/availability.model';
import { Accommodation } from 'src/app/model/accommodation.model';
import {FormControl, Validators, FormGroup, FormBuilder} from '@angular/forms';

@Component({
  selector: 'app-create-availability',
  templateUrl: './create-availability.component.html',
  styleUrls: ['./create-availability.component.css']
})
export class CreateAvailabilityComponent implements OnInit {

  constructor(
    private router: Router,
    private accommodationService: AccommodationService, 
    private userService : UserService,
    private route: ActivatedRoute,
  ) { }

  title = 'Create Availability';
  availability = new Availability();
  submitted = false;

  accommodation: Accommodation= new Accommodation(); 

  departureDateDateForm: Date = new Date();
  arrivalDateDateForm: Date = new Date();

  // selections: any[] =[]; 
  // priceSelection: string[] = ['Per Person', 'Per Accommodation'];
  // selected: string = '';

  isLoggedIn: boolean = false;
  isHost: boolean = false;
  isGuest: boolean = false;
  user: User = new User();
  //priceSelection: string = ''; 
  //startDate: string = ''
  accommodationId: string = ''
  price: number = 1
  startDate: Date = new Date()
  endDate: Date = new Date()
  priceSelection: number = 0


  ngOnInit(): void {
    // this.availability.accommodationId = this.router.params['id']; 

    this.route.params.subscribe((params: Params) => {
      this.accommodationService.getAccommodationById(params['id']).subscribe(res => {
        console.log(params['id'])
        this.accommodation = res.accommodation;
        console.log(res);
      })
    });


    let userRole = this.userService.getLoggedInUserRole()
    let userEmail = this.userService.getLoggedInUserEmail()
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

  onSubmit(){    
    console.log(this.startDate); 
    console.log(this.endDate);  
    var accommodation = {
      accommodationId : this.accommodation.id,
      availability : {
        startDate : this.startDate, 
        endDate : this.endDate,
        price : this.availability.price,
        priceSelection : Number(this.priceSelection)
      }
    }
      
    this.accommodationService.createAvailability(accommodation).subscribe(
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
