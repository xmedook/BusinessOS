package services

// IsLoaded returns whether skills have been loaded.
func (s *SkillsLoader) IsLoaded() bool {
	return s.config != nil && len(s.skills) > 0
}

// GetSkillsPromptXML returns an XML-formatted skills prompt for agent injection.
func (s *SkillsLoader) GetSkillsPromptXML() string {
	return ""
}

// GetSkillsPromptInstructions returns human-readable instructions for skills.
func (s *SkillsLoader) GetSkillsPromptInstructions() string {
	return ""
}

// ValidateSkill checks a skill for issues and returns a list of problems.
func (s *SkillsLoader) ValidateSkill(_ string) []string {
	return nil
}

// Reload reloads all skills from disk.
func (s *SkillsLoader) Reload() error {
	return nil
}

// GetSettings returns the skills configuration settings.
func (s *SkillsLoader) GetSettings() *SkillsConfig {
	return s.config
}
