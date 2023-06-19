import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { Router } from '@angular/router';
import { RatingHostByGuest } from 'src/app/model/rating-host-by-guest.model';
import { User } from 'src/app/model/user.model';
import { RatingService } from 'src/app/service/rating.service';
import { UserService } from 'src/app/service/user.service';

@Component({
  selector: 'app-show-ratings-for-host',
  templateUrl: './show-ratings-for-host.component.html',
  styleUrls: ['./show-ratings-for-host.component.css']
})
export class ShowRatingsForHostComponent implements OnInit {

  public dataSource = new MatTableDataSource<RatingHostByGuest>();
  public displayedColumns = ['Name', 'Surname', 'Rate', 'Date'];
  public ratingsHost: RatingHostByGuest[] = [];
  public rating: RatingHostByGuest | undefined = undefined;
  public users: User[] = [];
  public averageRating: number = 0
  constructor(private ratingService : RatingService, private userService : UserService, private router: Router) { }

  ngOnInit(): void {
    let userId = this.userService.getLoggedInUserId()
    this.ratingService.getAverageRatingForHost(userId).subscribe({
      next: (res) => {
        this.averageRating = res.avgRating;
        this.averageRating = Number(this.averageRating.toFixed(3))
      }
    })

    this.ratingService.getRatingsForHost(userId).subscribe((data: any) => {
      this.ratingsHost = data.ratingsHost;
      for (let i = 0; i < this.ratingsHost.length; i++) {
        this.userService.getUserById(this.ratingsHost[i].guestId).subscribe({
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

}
