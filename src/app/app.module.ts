import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';

import { MatButtonModule } from '@angular/material';

import { AppComponent } from './app.component';
import { NavComponent } from './header/nav/nav.component';
import { SidebarComponent } from './main/sidebar/sidebar.component';
import { Ch1Component } from './views/basic/ch1/ch1.component';
import { Ch2Component } from './views/basic/ch2/ch2.component';


@NgModule({
  declarations: [
    AppComponent,
    NavComponent,
    SidebarComponent,
    Ch1Component,
    Ch2Component
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatButtonModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
