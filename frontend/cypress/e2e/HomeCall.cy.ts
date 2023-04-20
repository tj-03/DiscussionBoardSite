describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/homepage')

    cy.contains('User').should('exist')
  })
})