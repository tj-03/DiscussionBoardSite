import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AboutteamComponent } from './aboutteam.component';

describe('AboutteamComponent', () => {
  let component: AboutteamComponent;
  let fixture: ComponentFixture<AboutteamComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AboutteamComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AboutteamComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
