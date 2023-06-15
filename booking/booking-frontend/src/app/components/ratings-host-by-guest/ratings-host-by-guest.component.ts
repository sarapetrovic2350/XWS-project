import { Component, OnInit } from '@angular/core';
import {MatTableDataSource} from "@angular/material/table";
import {User} from "../../model/user.model";
import {RatingHost} from "../../model/rating-host.model";
import {RatingService} from "../../service/rating.service";
import {UserService} from "../../service/user.service";
import {RatingHostByGuest} from "../../model/rating-host-by-guest.model";
import Swal from "sweetalert2";
import {Router} from "@angular/router";

@Component({
  selector: 'app-ratings-host-by-guest',
  templateUrl: './ratings-host-by-guest.component.html',
  styleUrls: ['./ratings-host-by-guest.component.css']
})
export class RatingsHostByGuestComponent implements OnInit {

  public dataSource = new MatTableDataSource<RatingHostByGuest>();
  public displayedColumns = ['Name', 'Surname', 'Rate', 'Date', 'edit', 'delete'];
  public ratingsHost: RatingHostByGuest[] = [];
  public rating: RatingHostByGuest | undefined = undefined;
  public users: User[] = [];
  constructor(private ratingService : RatingService, private userService : UserService, private router: Router) { }

  ngOnInit(): void {
    let userId = this.userService.getLoggedInUserId()
    this.ratingService.getRatingsHostByGuestId(userId).subscribe((data: any) => {
      this.ratingsHost = data.ratingsHost;
      for (let i = 0; i < this.ratingsHost.length; i++) {
        this.userService.getUserById(this.ratingsHost[i].hostId).subscribe({
          next: (res) => {
            this.ratingsHost[i].hostName = res.user.firstName;
            this.ratingsHost[i].hostSurname = res.user.lastName;
          }
        })
      }
      console.log(this.ratingsHost);
      this.dataSource.data = this.ratingsHost;
    })
  }
  edit(id: string) {
    this.router.navigate(['edit-rating-host/' + id]);
  }
  delete(id: string) {
    this.ratingService.DeleteRatingForHost(id).subscribe(
      {
        next: (res) => {
          this.router.navigate(['/ratings-host-by-guest']);
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: 'Successfully deleted rating for host!',
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
