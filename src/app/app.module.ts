import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';

import { MatButtonModule } from '@angular/material';

import { AppComponent } from './app.component';
import { NavComponent } from './header/nav/nav.component';
import { SidebarComponent } from './main/sidebar/sidebar.component';
import { Ch1Component } from './views/basic/ch1/ch1.component';
import { Ch2Component } from './views/basic/ch2/ch2.component';
import { EditorComponent } from './main/editor/editor.component';
import { MainService } from './services/main.service';
import { TestresultComponent } from './main/testresult/testresult.component';


@NgModule({
  declarations: [
    AppComponent,
    NavComponent,
    SidebarComponent,
    Ch1Component,
    Ch2Component,
    EditorComponent,
    TestresultComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MatButtonModule,
    HttpClientModule,
    AppRoutingModule,
    FormsModule
  ],
  providers: [
    MainService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
