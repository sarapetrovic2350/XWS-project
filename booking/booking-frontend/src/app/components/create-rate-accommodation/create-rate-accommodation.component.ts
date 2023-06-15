import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { Accommodation } from 'src/app/model/accommodation.model';
import { RatingAccommodation } from 'src/app/model/rating-accommodation.model';
import { RatingHost } from 'src/app/model/rating-host.model';
import { User } from 'src/app/model/user.model';
import { AccommodationService } from 'src/app/service/accommodation.service';
import { RatingService } from 'src/app/service/rating.service';
import { UserService } from 'src/app/service/user.service';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-create-rate-accommodation',
  templateUrl: './create-rate-accommodation.component.html',
  styleUrls: ['./create-rate-accommodation.component.css']
})
export class CreateRateAccommodationComponent implements OnInit {

  constructor(
    private router: Router,
    private accommodationService : AccommodationService,
    private route: ActivatedRoute,
    private userService: UserService,
    private ratinService: RatingService
  ) { }

  title = 'Rate accommodation';
  submitted = false;  

  accommodation: Accommodation = new Accommodation();
  user: User = new User();
  ratingAccommodation: RatingAccommodation = new RatingAccommodation();

  ngOnInit(): void {

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
      this.user = res.user;
      console.log(res.user.id)
    })


  }

  onSubmit(){

    this.ratingAccommodation.guestId = this.user.id;
    this.ratingAccommodation.accommodationId = this.accommodation.id;
    console.log(this.ratingAccommodation)

    this.ratinService.createRatingForAccommodation(this.ratingAccommodation).subscribe(
      {
        next: (res) => {
          //this.router.navigate(['/show-host-accommodations']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully rated host!',
          })

        },
        error: (e) => {
          this.submitted = false;
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Guest did not stay at accommodation he wants to rate.',
          })

        }

      });

  }

}
