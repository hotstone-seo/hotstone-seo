describe('rules_pagination', () => {
  beforeEach(() => {
    cy.fixture('cookies.json').then((cookies) => {
      cy.setCookie('token', cookies.token);
    });

    cy.server();

    cy.fixture('rules/get_rules_resp_offset0_limit10_total5.json').as('get_rules_resp_offset0_limit10_total5');
    cy.fixture('rules/get_rules_resp_offset0_limit10_total8.json').as('get_rules_resp_offset0_limit10_total8');
    cy.fixture('rules/get_rules_resp_offset5_limit10_total3.json').as('get_rules_resp_offset5_limit10_total3');
  });

  context('pagination', () => {
    it('GIVEN total <= pageSize THEN show only one page', () => {
      cy.visit('/rules');

      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total5')
        .as('get_rules_offset0_limit10');

      cy.wait('@get_rules_offset0_limit10');

      cy.get('.ant-pagination-next > a').should('have.attr', 'disabled');
    });

    it('GIVEN total > pageSize THEN show two pages', () => {
      cy.visit('/rules');

      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total8')
        .as('get_rules_offset0_limit10');

      cy.wait('@get_rules_offset0_limit10');

      cy.get('.ant-pagination-next > a').should('not.have.attr', 'disabled');
    });

    it('GIVEN two pages AND click page 2 THEN load page 2', () => {
      cy.visit('/rules');

      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total8')
        .as('get_rules_offset0_limit10');

      cy.wait('@get_rules_offset0_limit10');

      cy.route('GET', '/api/rules?_offset=5&_limit=10', '@get_rules_resp_offset5_limit10_total3')
        .as('get_rules_offset5_limit10');

      cy.get('.ant-pagination-next > a').click();

      cy.wait('@get_rules_offset5_limit10');

      cy.get('.ant-pagination-next > a').should('have.attr', 'disabled');
      cy.get('[data-row-key]').should('have.length', 3);
    });
  });
});
