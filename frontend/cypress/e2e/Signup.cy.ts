describe('Signup.cy.ts', () => {
  it('ensures sign up process works correctly', () => {
    cy.visit('http://localhost:4200')

    cy.contains('Log in').click()

    cy.contains('Sign up').click()

    cy.url().should('include', '/signup')

    cy.contains('Email').should('exist')
    cy.contains('Username').should('exist')
    cy.contains('Password').should('exist')
  })
})