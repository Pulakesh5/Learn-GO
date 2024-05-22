initial_balance = 0
per_month_contribution = 51000
years = 1
period = 11
final_balance = 0
rate = 0.05 #monthly 

total_paid = 0

# Phase 1: Contributing for the first 5 years
for years in range(1, 6):
    for i in range((years - 1) * 12, 12 * years):
        final_balance = (final_balance + per_month_contribution) * (1 + rate)
        total_paid += per_month_contribution
    print(f"Balance after {years} years: ₹{final_balance:.2f} (Total paid: ₹{total_paid:.2f})")

# Phase 2: No contributions for the next 5 years (only interest accumulation)

for years in range(6, period):
    for i in range(12):  # 12 months of interest calculation each year
        final_balance *= (1 + rate)  # Calculate interest only
    print(f"Balance after {years} years: ₹{final_balance:.2f} (Total paid: ₹{total_paid:.2f})")  # Total paid remains the same
