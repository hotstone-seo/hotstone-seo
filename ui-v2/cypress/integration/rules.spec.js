describe('/rules', () => {
  beforeEach(() => {
    cy.fixture('cookies.json').then((cookies) => {
      cy.setCookie('token', cookies.token);
    });

    cy.server();

    cy.fixture('rules/_offset0_limit10_total5.json').as('data_rules_offset0_limit10_total5');
    cy.fixture('rules/_offset0_limit10_total8.json').as('data_rules_offset0_limit10_total8');
    cy.fixture('rules/_offset5_limit10_total3.json').as('data_rules_offset5_limit10_total3');
  });

  context('pagination', () => {
    it('GIVEN total <= pageSize THEN show only one page', () => {
      cy.visit('/rules');

      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@data_rules_offset0_limit10_total5')
        .as('route_rules_offset0_limit10');

      cy.wait('@route_rules_offset0_limit10');

      cy.get('.ant-pagination-next > a').should('have.attr', 'disabled');
    });

    it('GIVEN total > pageSize THEN show two pages', () => {
      cy.visit('/rules');

      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@data_rules_offset0_limit10_total8')
        .as('route_rules_offset0_limit10');

      cy.wait('@route_rules_offset0_limit10');

      cy.get('.ant-pagination-next > a').should('not.have.attr', 'disabled');
    });

    it('GIVEN two pages AND click page 2 THEN load page 2', () => {
      cy.visit('/rules');

      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@data_rules_offset0_limit10_total8')
        .as('route_rules_offset0_limit10');

      cy.wait('@route_rules_offset0_limit10');

      cy.route('GET', '/api/rules?_offset=5&_limit=10', '@data_rules_offset5_limit10_total3')
        .as('route_rules_offset5_limit10');

      cy.get('.ant-pagination-next > a').click();

      cy.wait('@route_rules_offset5_limit10');

      cy.get('.ant-pagination-next > a').should('have.attr', 'disabled');
      cy.get('[data-row-key]').should('have.length', 3);
    });
  });
});
