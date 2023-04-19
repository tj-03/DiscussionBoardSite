describe('Logo', () => {
    it('ensures "Heap Underflow" redirects to /homepage', () => {
      cy.visit('http://localhost:4200')
  
      cy.contains('Heap Underflow').click()
  
      cy.url().should('include', '/homepage')
    })
  })