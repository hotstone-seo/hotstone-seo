describe('Tags', () => {
  beforeEach(() => {
    cy.fixture('cookies.json').then((cookies) => {
      cy.setCookie('token', cookies.token);
    });

    cy.server();

    cy.fixture('tags/get_rules_resp_offset0_limit10_total5.json').as('get_rules_resp_offset0_limit10_total5');
    cy.fixture('tags/get_rules_5_resp.json').as('get_rules_5_resp');

    const resizeObserverLoopErrRe = /^ResizeObserver loop limit exceeded/;
    Cypress.on('uncaught:exception', (err) => {
      if (resizeObserverLoopErrRe.test(err.message)) {
        return false;
      }
    });
  });

  context('Add', () => {
    it('new Title Tag', () => {
      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total5')
        .as('get_rules_offset0_limit10');
      cy.route('GET', '/api/rules/5', '@get_rules_5_resp')
        .as('get_rules_5');
      cy.route('POST', '/api/tags', {})
        .as('post_tags');

      cy.visit('/rules');
      cy.wait('@get_rules_offset0_limit10');

      cy.get('[data-row-key=5] > :nth-child(2) > [data-testid=btn-detail]').click();

      cy.url().should('eq', `${Cypress.config().baseUrl}/rules/5`);
      cy.wait('@get_rules_5');

      cy.get('[data-testid=btn-new-tag]').click();

      cy.get('[data-testid="select-type"]').type('Title{enter}');
      cy.get('[data-testid="select-locale"]').type('id_ID{enter}');
      cy.get('[data-testid="input-title"]').type('Foo Title Tag');

      cy.get('[data-testid="text-preview-tag"]').should('have.text', '<title>Foo Title Tag</title>');

      cy.get('[data-testid="btn-save-tag"]').click();

      cy.get('@post_tags')
        .its('request.body')
        .should('deep.equal', {
          rule_id: 5,
          type: 'title',
          locale: 'id_ID',
          value: 'Foo Title Tag',
        });
    });

    it('new Meta Tag', () => {
      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total5')
        .as('get_rules_offset0_limit10');
      cy.route('GET', '/api/rules/5', '@get_rules_5_resp')
        .as('get_rules_5');
      cy.route('POST', '/api/tags', {})
        .as('post_tags');

      cy.visit('/rules');
      cy.wait('@get_rules_offset0_limit10');

      cy.get('[data-row-key=5] > :nth-child(2) > [data-testid=btn-detail]').click();

      cy.url().should('eq', `${Cypress.config().baseUrl}/rules/5`);
      cy.wait('@get_rules_5');

      cy.get('[data-testid=btn-new-tag]').click();

      cy.get('[data-testid="select-type"]').type('Meta{enter}');
      cy.get('[data-testid="select-locale"]').type('id_ID{enter}');
      cy.get('[data-testid="input-name"]').type('description');
      cy.get('[data-testid="input-content"]').type('Foo Content');

      cy.get('[data-testid="text-preview-tag"]').should('have.text', '<meta name="description" content="Foo Content"/>');

      cy.get('[data-testid="btn-save-tag"]').click();

      cy.get('@post_tags')
        .its('request.body')
        .should('deep.equal', {
          rule_id: 5,
          type: 'meta',
          locale: 'id_ID',
          attributes: {
            name: 'description',
            content: 'Foo Content',
          },
        });
    });
  });
});
