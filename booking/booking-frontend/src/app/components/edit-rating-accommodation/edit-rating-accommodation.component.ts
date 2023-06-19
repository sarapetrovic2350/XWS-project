import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Params, Router} from "@angular/router";
import {RatingService} from "../../service/rating.service";
import {UserService} from "../../service/user.service";
import {RatingHost} from "../../model/rating-host.model";
import {RatingAccommodation} from "../../model/rating-accommodation.model";
import {AccommodationService} from "../../service/accommodation.service";
import Swal from "sweetalert2";

@Component({
  selector: 'app-edit-rating-accommodation',
  templateUrl: './edit-rating-accommodation.component.html',
  styleUrls: ['./edit-rating-accommodation.component.css']
})
export class EditRatingAccommodationComponent implements OnInit {

  constructor(private ratingService: RatingService,private accommodationService: AccommodationService, private router: Router, private route: ActivatedRoute) { }
  ratingAccommodation: RatingAccommodation = new RatingAccommodation();
  accommodationName: string = "";
  title = 'Edit rating for accommodation';
  submitted = false;
  stars: number[] = [1, 2, 3, 4, 5];
  selectedValue: number = 0;
  ngOnInit(): void {
    this.route.params.subscribe((params: Params) => {
      this.ratingService.getRatingAccommodationById(params['id']).subscribe(res => {
        console.log(params['id'])
        this.ratingAccommodation = res.ratingAccommodation;
        this.selectedValue = this.ratingAccommodation.rate;
        console.log(this.ratingAccommodation);
        this.accommodationService.getAccommodationById(this.ratingAccommodation.accommodationId).subscribe(res => {
          console.log(res)
          this.accommodationName = res.accommodation.name;
        })
      })
    });
  }
  onSubmit() {
    console.log(this.ratingAccommodation)
    this.ratingAccommodation.rate = this.selectedValue;
    this.ratingService.updateRatingForAccommodation(this.ratingAccommodation).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/ratings-accommodation-by-guest']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully updated rate for accommodation!',
          })
        },
        error: (e) => {
          console.log(e);
          Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Something went wrong.',
          })

        }

      });
  }
  countStar(star: number) {
    console.log('Value of star', star);
    this.selectedValue = star;
    console.log(this.selectedValue)
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
