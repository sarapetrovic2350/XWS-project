import { Component, OnInit } from '@angular/core';
import {MatTableDataSource} from "@angular/material/table";
import {RatingHostByGuest} from "../../model/rating-host-by-guest.model";
import {User} from "../../model/user.model";
import {RatingService} from "../../service/rating.service";
import {UserService} from "../../service/user.service";
import {Router} from "@angular/router";
import Swal from "sweetalert2";
import {AccommodationService} from "../../service/accommodation.service";
import {RatingAccommodationByGuest} from "../../model/rating-accommodation-by-guest.model";

@Component({
  selector: 'app-ratings-accommodation-by-guest',
  templateUrl: './ratings-accommodation-by-guest.component.html',
  styleUrls: ['./ratings-accommodation-by-guest.component.css']
})
export class RatingsAccommodationByGuestComponent implements OnInit {
  public dataSource = new MatTableDataSource<RatingAccommodationByGuest>();
  public displayedColumns = ['Name', 'Rate', 'Date', 'edit', 'delete'];
  public ratingsAccommodation: RatingAccommodationByGuest[] = [];
  public rating: RatingAccommodationByGuest | undefined = undefined;
  public users: User[] = [];
  constructor(private userService: UserService, private ratingService : RatingService, private accommodationService : AccommodationService, private router: Router) { }

  ngOnInit(): void {
    let userId = this.userService.getLoggedInUserId()
    this.ratingService.getRatingsAccommodationByGuestId(userId).subscribe((data: any) => {
      this.ratingsAccommodation = data.ratingsAccommodation;
      for (let i = 0; i < this.ratingsAccommodation.length; i++) {
        this.accommodationService.getAccommodationById(this.ratingsAccommodation[i].accommodationId).subscribe({
          next: (res) => {
            this.ratingsAccommodation[i].accommodationName = res.accommodation.name;
          }
        })
      }
      console.log(this.ratingsAccommodation);
      this.dataSource.data = this.ratingsAccommodation;
    })
  }
  edit(id: string) {
    this.router.navigate(['edit-rating-accommodation/' + id]);
  }
  delete(id: string) {
    this.ratingService.DeleteRatingForAccommodation(id).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/ratings-accommodation-by-guest']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully deleted rating for accommodation!',
          })
          window.location.reload();
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

}
