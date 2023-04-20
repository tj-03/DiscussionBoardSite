describe('AboutTeam', () => {
  it('ensures "About the Team" redirects to /about', () => {
    cy.visit('http://localhost:4200')

    cy.contains('About the Team').click()

    cy.url().should('include', '/about')
  })
})