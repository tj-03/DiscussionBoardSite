describe('CreatePost', () => {
  it('ensures "Create Post" redirects to /createpost', () => {
    cy.visit('http://localhost:4200')

    cy.contains('Create a Post').click()

    cy.url().should('include', '/createpost')
  })
})