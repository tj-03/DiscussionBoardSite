import { StepperComponent } from './stepper.component'

describe('aboutteam', () => {
  it('ensures "About the Team" redirects to /aboutteam', () => {
    cy.visit('http://localhost:4200')

    cy.contains('About the Team').click()

    cy.url().should('include', '/aboutteam')
  })
})
