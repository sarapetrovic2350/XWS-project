import { Component, OnInit } from '@angular/core';
import {RatingService} from "../../service/rating.service";
import {UserService} from "../../service/user.service";
import {Router} from "@angular/router";
import {MatTableDataSource} from "@angular/material/table";
import {RatingHostByGuest} from "../../model/rating-host-by-guest.model";
import {RatingAccommodationForHost} from "../../model/rating-accommodation-for-host.model";
import {User} from "../../model/user.model";
import {AccommodationService} from "../../service/accommodation.service";

@Component({
  selector: 'app-show-ratings-accommodations-host',
  templateUrl: './show-ratings-accommodations-host.component.html',
  styleUrls: ['./show-ratings-accommodations-host.component.css']
})
export class ShowRatingsAccommodationsHostComponent implements OnInit {
  public dataSource = new MatTableDataSource<RatingAccommodationForHost>();
  public displayedColumns = ['Name', 'Surname', 'Rate', 'Date'];
  public ratingsAccommodation: RatingAccommodationForHost[] = [];
  public rating: RatingAccommodationForHost | undefined = undefined;
  public averageRating: number = 0
  public totalRating: number = 0
  constructor(
    private accommodationService: AccommodationService,
    private ratingService : RatingService,
    private userService : UserService,
    private router: Router
  ) { }

  ngOnInit(): void {
    let userId = this.userService.getLoggedInUserId()

    this.ratingService.getRatingsAccommodationsForHost(userId).subscribe((data: any) => {
      this.ratingsAccommodation = data.ratingsAccommodation;
      for (let i = 0; i < this.ratingsAccommodation.length; i++) {
        this.totalRating += this.ratingsAccommodation[i].rate
        this.userService.getUserById(this.ratingsAccommodation[i].guestId).subscribe({
          next: (res) => {
            this.ratingsAccommodation[i].guestName = res.user.firstName;
            this.ratingsAccommodation[i].guestSurname = res.user.lastName;
          }
        })
        this.accommodationService.getAccommodationById(this.ratingsAccommodation[i].accommodationId).subscribe({
          next: (res) => {
            this.ratingsAccommodation[i].accommodationName = res.accommodation.name;
          }
        })
      }
      this.averageRating = this.totalRating / this.ratingsAccommodation.length;
      this.averageRating = Number(this.averageRating.toFixed(3))
      console.log(this.ratingsAccommodation);
      this.dataSource.data = this.ratingsAccommodation;
    })

  }
}
