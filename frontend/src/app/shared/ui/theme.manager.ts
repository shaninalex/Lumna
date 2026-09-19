import type { BehaviorSubject } from 'rxjs';

export type Theme = 'light' | 'dark' | 'auto';

export class ThemeManager {
    private readonly storageKey = 'theme';
    private mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
    private mediaQueryListener = () => {
        if (this.theme$.value === 'auto') {
            this.updateDomTheme('auto');
        }
    };

    constructor(private theme$: BehaviorSubject<Theme>) {}

    public init() {
        const saved = this.getStoredTheme();
        const theme: Theme = saved ?? 'auto';
        this.theme$.next(theme);
        this.updateDomTheme(theme);

        if (this.mediaQuery.addEventListener) {
            this.mediaQuery.addEventListener('change', this.mediaQueryListener);
        } else {
            this.mediaQuery.addListener(this.mediaQueryListener);
        }
    }

    private getStoredTheme(): Theme | null {
        try {
            const stored = localStorage.getItem(this.storageKey);
            if (stored === 'light' || stored === 'dark' || stored === 'auto') {
                return stored;
            }
        } catch {
            return null;
        }
        return null;
    }

    private getSystemTheme(): 'light' | 'dark' {
        return this.mediaQuery.matches ? 'dark' : 'light';
    }

    private updateDomTheme(theme: Theme) {
        const resolvedTheme = theme === 'auto' ? this.getSystemTheme() : theme;
        document.documentElement.setAttribute('data-bs-theme', resolvedTheme);
    }

    public applyTheme(theme: Theme) {
        this.theme$.next(theme);
        this.updateDomTheme(theme);
        try {
            localStorage.setItem(this.storageKey, theme);
        } catch {
            // Ignore storage errors in restricted contexts
        }
    }

    public toggleTheme() {
        const next: Theme =
            this.theme$.value === 'light'
                ? 'dark'
                : this.theme$.value === 'dark'
                  ? 'auto'
                  : 'light';
        this.applyTheme(next);
    }
}
