package android
type Product_variables struct {
	Allows_invalid_pthread struct {
		Cppflags []string
	}
	Egl_needs_handle struct {
		Cppflags []string
	}
	Egl_workaround_bug_10194508 struct {
		Cppflags []string
	}
	Exynos4_enhancements struct {
		Cflags []string
		Cppflags []string
	}
	Needs_text_relocations struct {
		Cppflags []string
	}
}

type ProductVariables struct {
	Allows_invalid_pthread    *bool `json:",omitempty"`
	Egl_needs_handle    *bool `json:",omitempty"`
	Egl_workaround_bug_10194508    *bool `json:",omitempty"`
	Exynos4_enhancements    *bool `json:",omitempty"`
	Needs_text_relocations  *bool `json:",omitempty"`
}
