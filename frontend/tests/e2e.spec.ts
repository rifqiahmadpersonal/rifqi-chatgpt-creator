import { test, expect } from '@playwright/test';

const BASE_URL = process.env.FRONTEND_URL || 'http://localhost:3000';

test.describe('Home Page', () => {
  test('should load home page', async ({ page }) => {
    await page.goto(`${BASE_URL}/`);
    await expect(page).toHaveTitle(/ChatGPT/i);
    await expect(page.locator('h1')).toBeVisible();
  });
});

test.describe('Accounts Page', () => {
  test('should load accounts page', async ({ page }) => {
    await page.goto(`${BASE_URL}/accounts`);
    await expect(page.locator('h1')).toContainText('Accounts');
  });
});

test.describe('Batch Jobs Page', () => {
  test('should load batch jobs page', async ({ page }) => {
    await page.goto(`${BASE_URL}/batch-jobs`);
    await expect(page.locator('h1')).toContainText('Batch');
  });
});

test.describe('Blacklisted Domains Page', () => {
  test('should load blacklisted domains page', async ({ page }) => {
    await page.goto(`${BASE_URL}/blacklisted-domains`);
    await expect(page.locator('h1')).toContainText('Blacklisted');
  });
});

test.describe('Configuration Page', () => {
  test('should load configuration page', async ({ page }) => {
    await page.goto(`${BASE_URL}/configuration`);
    await expect(page.locator('h1')).toContainText('Configuration');
  });
});

test.describe('Email Domains Page', () => {
  test('should load email domains page', async ({ page }) => {
    await page.goto(`${BASE_URL}/email-domains`);
    await expect(page.locator('h1')).toContainText('Email');
  });
});

test.describe('Stats Page', () => {
  test('should load stats page', async ({ page }) => {
    await page.goto(`${BASE_URL}/stats`);
    await expect(page.locator('h1')).toContainText('Dashboard');
  });
});