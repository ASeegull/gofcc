import { Component, OnInit, ViewChild } from '@angular/core';
import { MainService } from '../../services/main.service';

@Component({
  selector: 'app-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css']
})
export class SidebarComponent implements OnInit {

  constructor(public mainService: MainService) { }

  ngOnInit() {
  }

  run() {
    this.mainService.runTests(this.mainService.solution);
  }

  format() {
    this.mainService.fmt(this.mainService.solution);
  }
}
