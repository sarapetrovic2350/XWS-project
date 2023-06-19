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
    private ratingService: RatingService
  ) { }

  title = 'Rate accommodation';
  submitted = false;

  accommodation: Accommodation = new Accommodation();
  user: User = new User();
  ratingAccommodation: RatingAccommodation = new RatingAccommodation();
  stars: number[] = [1, 2, 3, 4, 5];
  selectedValue: number = 0;

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
    this.ratingAccommodation.rate = this.selectedValue;
    console.log(this.ratingAccommodation)

    this.ratingService.createRatingForAccommodation(this.ratingAccommodation).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/ratings-accommodation-by-guest']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully rated accommodation!',
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

  countStar(star: number) {
    console.log('Value of star', star);
    this.selectedValue = star;
  }
  addClass(star: number) {
    let ab = "";
    for (let i = 0; i < star; i++) {
      ab = "starId" + i;
      // @ts-ignore
      document.getElementById(ab).classList.add("selected");
    }
  }
  removeClass(star: number) {
    let ab = "";
    for (let i = star-1; i >= this.selectedValue; i--) {
      ab = "starId" + i;
      // @ts-ignore
      document.getElementById(ab).classList.remove("selected");
    }
  }

}
