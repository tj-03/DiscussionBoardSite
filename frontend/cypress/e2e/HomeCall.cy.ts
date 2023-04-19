describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/homepage')

    cy.contains('Post ID').should('exist')
  })
})